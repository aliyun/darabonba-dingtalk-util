<?php

namespace AlibabaCloud\DingUtil;

/**
 * This is for DingUtil.
 */
class DingUtilClient
{
    /**
     * @return string
     */
    public static function getTimestamp()
    {
        return (string) (int) (microtime(true) * 1000);
    }

    /**
     * @param array $res
     *
     * @return bool
     */
    public static function hasError($res)
    {
        if (isset($res['error_response'])) {
            return true;
        }
        if (!isset($res['errcode'])) {
            return false;
        }
        if (0 == (int) ($res['errcode'])) {
            return false;
        }

        return true;
    }

    /**
     * @param string $accessSecret
     * @param string $canonicalString
     *
     * @return string
     */
    public static function computeSignature($accessSecret, $canonicalString)
    {
        return base64_encode(hash_hmac('sha256', $canonicalString, $accessSecret, true));
    }

    /**
     * @param string $timestamp
     * @param string $suiteTicket
     *
     * @return string
     */
    public static function getCanonicalStringForIsv($timestamp, $suiteTicket)
    {
        $res = (string) $timestamp;
        if (null !== $suiteTicket) {
            $res .= "\n" . $suiteTicket;
        }

        return $res;
    }
}
