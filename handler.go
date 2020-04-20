package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type WebHookHandler func(eventname string, payload *GitHubPayload, req *http.Request) error

func Handler(secret string, fn WebHookHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		event := c.GetHeader("x-github-event")
		delivery := c.GetHeader("x-github-delivery")
		signature := c.GetHeader("x-hub-signature")
		contentType := c.GetHeader("Content-Type")

		//log.Printf("event:%s, delivery:%s, sign:%s \n", event, delivery, signature)
		// Utility funcs
		_fail := func(err error) {
			fail(c, event, err)
		}
		_succeed := func() {
			succeed(c, event)
		}

		// Ensure headers are all there
		if event == "" || delivery == "" {
			_fail(fmt.Errorf("Missing x-github-* and x-hub-* headers"))
			return
		}

		// No secret provided to github
		if signature == "" && secret != "" {
			_fail(fmt.Errorf("GitHub isn't providing a signature, whilst a secret is being used (please give github's webhook the secret)"))
			return
		}

		// Read body
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			_fail(err)
			return
		}

		// Validate payload (only when secret is provided)
		if secret != "" {
			if err := validePayloadSignature(secret, signature, body); err != nil {
				// Valied validation
				_fail(err)
				return
			}
		}
		// Get payload. from github data is encode. fix by 2020.04.20
		fmt.Println(string(body))
		var payloadData []byte
		if contentType == "application/x-www-form-urlencoded" {
			payload_split := strings.SplitN(string(body), "=", 2)
			payloadUrlDecode, err := url.QueryUnescape(payload_split[1])
			if err != nil {
				_fail(err)
				return
			}
			payloadData = []byte(payloadUrlDecode)
		} else if contentType == "application/json" {

		}

		payload := GitHubPayload{}
		if err := json.Unmarshal(payloadData, &payload); err != nil {
			_fail(fmt.Errorf("Could not deserialize payload"))
			return
		}
		// Do something with payload
		if err := fn(event, &payload, c.Request); err == nil {
			_succeed()
		} else {
			_fail(err)
		}
	}
}

func validePayloadSignature(secret, signatureHeader string, body []byte) error {
	// Check header is valid
	signature_parts := strings.SplitN(signatureHeader, "=", 2)
	if len(signature_parts) != 2 {
		return fmt.Errorf("Invalid signature header: '%s' does not contain two parts (hash type and hash)", signatureHeader)
	}

	// Ensure secret is a sha1 hash
	signature_type := signature_parts[0]
	signature_hash := signature_parts[1]
	if signature_type != "sha1" {
		return fmt.Errorf("Signature should be a 'sha1' hash not '%s'", signature_type)
	}

	// Check that payload came from github
	// skip check if empty secret provided
	if !IsValidPayload(secret, signature_hash, body) {
		return fmt.Errorf("Payload did not come from GitHub")
	}

	return nil
}

func succeed(c *gin.Context, event string) {
	c.JSON(200, PayloadPong{
		Ok:    true,
		Event: event,
	})
}

func fail(c *gin.Context, event string, err error) {
	log.Printf("event:%s, err:%s", event, err.Error())
	c.JSON(500, PayloadPong{
		Ok:    false,
		Event: event,
		Error: err.Error(),
	})
}
