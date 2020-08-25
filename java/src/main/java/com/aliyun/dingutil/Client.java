// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingutil;


import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.text.SimpleDateFormat;
import java.util.Base64;
import java.util.Date;
import java.util.SimpleTimeZone;


public class Client {

    public static String getTimestamp() throws Exception {
        SimpleDateFormat df = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss'Z'");
        df.setTimeZone(new SimpleTimeZone(0, "UTC"));
        return df.format(new Date());
    }

    public static Boolean hasError(java.util.Map<String, ?> res) throws Exception {
        if (null == res) {
            return true;
        }
        if (null != res.get("error_response")) {
            return true;
        }

        if (null == res.get("errcode")) {
            return false;
        }
        Object errCode = res.get("errcode");
        if (null == errCode && (int) errCode == 0) {
            return false;
        }
        return true;
    }

    public static String computeSignature(String accessSecret, String canonicalString) throws Exception {
        byte[] signData;
        Mac mac = Mac.getInstance("HMACSHA256");
        mac.init(new SecretKeySpec(accessSecret.getBytes("UTF-8"), "HMACSHA256"));
        signData = mac.doFinal(canonicalString.getBytes("UTF-8"));
        String signedStr = Base64.getEncoder().encodeToString(signData);
        return signedStr;
    }

    public static String getCanonicalStringForIsv(String timestamp, String suiteTicket) throws Exception {
        String result = timestamp;
        if (suiteTicket != null) {
            result = "\n" + suiteTicket;
        }
        return result;
    }
}
