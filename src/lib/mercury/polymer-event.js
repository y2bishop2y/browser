// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

var BaseEvent = require('mercury').BaseEvent;

// export extended BaseEvent!
module.exports = new BaseEvent(internalPolymerEvent);

/*
 * This allows us to attach a single event listener at the root of a VDOM tree
 * and receive polymer events from any item in the tree
 */

function internalPolymerEvent(ev, broadcast) {
    this.data.polymerDetail = ev._rawEvent.detail;
    this.data.target = ev._rawEvent.target;
    this.data.rawEvent = ev._rawEvent;
    broadcast(this.data);
}
