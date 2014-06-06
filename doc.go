// Copyright 2014 Vic Demuzere
//
// Use of this source code is governed by the MIT license.

// Package irc provides tools for communicating with IRC servers.
//
// The Message and Prefix structs provide translation to and from raw IRC messages:
//
//    // Parse the IRC-encoded data and store the result in a new struct:
//    message := irc.ParseMessage(raw)
//
//    // Translate back to a raw IRC message string:
//    raw = message.String()
//
// Decoder and Encoder can be used to decode and encode messages in a stream:
//
//    // Create a decoder that reads from given io.Reader
//    dec := irc.NewDecoder(reader)
//
//    // Decode the next IRC message
//    message, err := dec.Decode()
//
//    // Create an encoder that writes to given io.Writer
//    enc := irc.NewEncoder(writer)
//
//    // Send a message to the writer.
//    enc.Encode(message)
//
package irc
