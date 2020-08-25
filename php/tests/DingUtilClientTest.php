<?php

namespace AlibabaCloud\DingUtil\Tests;

use AlibabaCloud\DingUtil\DingUtilClient;
use PHPUnit\Framework\TestCase;

/**
 * @internal
 * @coversNothing
 */
class DingUtilClientTest extends TestCase
{
    public function testGetTimestamp()
    {
        $timestamp = DingUtilClient::getTimestamp();
        $this->assertTrue(is_numeric($timestamp) && 13 === \strlen($timestamp));
    }

    public function testHasError()
    {
        $res = ['error_response' => true];
        $this->assertTrue(DingUtilClient::hasError($res));

        $res = ['data' => []];
        $this->assertFalse(DingUtilClient::hasError($res));

        $res = ['errcode' => 0];
        $this->assertFalse(DingUtilClient::hasError($res));

        $res = ['errcode' => 500];
        $this->assertTrue(DingUtilClient::hasError($res));
    }

    public function testComputeSignature()
    {
        $this->assertEquals(
            'Ku8nmti/NjJ5kJzEyzXkpgNSuzwn4TAh3rxsBSaA/1E=',
            DingUtilClient::computeSignature('accessSecret', 'test')
        );
    }

    public function testGetCanonicalStringForIsv()
    {
        $timestamp = DingUtilClient::getTimestamp();
        $this->assertEquals(
            $timestamp,
            DingUtilClient::getCanonicalStringForIsv(DingUtilClient::getTimestamp(), null)
        );

        $this->assertEquals(
            $timestamp . "\nsuiteTicket",
            DingUtilClient::getCanonicalStringForIsv(DingUtilClient::getTimestamp(), 'suiteTicket')
        );
    }
}
