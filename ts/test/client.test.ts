import Client from "../src/client";

import * as $tea from "@alicloud/tea-typescript";
import assert from "assert";
import "mocha";

describe("ding-util", function () {
  it("Module should ok", function () {
    assert.ok(Client);
  });
});

describe("getTimestamp", function () {
  it("timestamp should ok", function () {
    const timestamp = Client.getTimestamp();
    assert.equal(typeof timestamp, "string");
    assert.equal(timestamp.length, 13);
  });
});

describe("testHasError", function () {
  it("test null", () => {
    const data = null;
    const res = Client.hasError(data);
    assert.equal(res, true);
  });

  it("test undefined", () => {
    const data = undefined;
    const res = Client.hasError(data);
    assert.equal(res, true);
  });

  it("test error_response", () => {
    const data = { error_response: true };
    const res = Client.hasError(data);
    assert.equal(res, true);
  });

  it("test errcode null", () => {
    const data = { errcode: null };
    const res = Client.hasError(data);
    assert.equal(res, false);
  });

  it("test errcode undefined", () => {
    const data = { errcode: undefined };
    const res = Client.hasError(data);
    assert.equal(res, false);
  });

  it("test errcode 0", () => {
    const data = { errcode: 0 };
    const res = Client.hasError(data);
    assert.equal(res, false);
  });

  it("test errcode 500", () => {
    const data = { errcode: 500 };
    const res = Client.hasError(data);
    assert.equal(res, true);
  });
});

describe("testComputeSignature", function () {
  it("computeSignature", () => {
    assert.equal(
      Client.computeSignature("accessSecret", "test"),
      "Ku8nmti/NjJ5kJzEyzXkpgNSuzwn4TAh3rxsBSaA/1E="
    );
  });
});

describe("getCanonicalStringForIsv", function () {
  it("getCanonicalStringForIsv", () => {
    const timestamp = Client.getTimestamp();
    assert.equal(
      Client.getCanonicalStringForIsv(Client.getTimestamp(), null),
      timestamp
    );
    assert.equal(
      Client.getCanonicalStringForIsv(Client.getTimestamp(), "suiteTicket"),
      `${timestamp}\nsuiteTicket`
    );
  });
});
