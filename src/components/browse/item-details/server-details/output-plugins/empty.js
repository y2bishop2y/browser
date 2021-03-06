// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

var h = require('mercury').h;

module.exports = {
  'shouldFormat': shouldFormat,
  'format': format
};

/*
 * The input is empty in various cases.
 */
function shouldFormat(input) {
  return input === undefined || input === null || input === '' ||
    (input instanceof Array && input.length === 0);
}

/*
 * Indicate that nothing was there.
 */
function format(input) {
  return h('span', '<no data>');
}