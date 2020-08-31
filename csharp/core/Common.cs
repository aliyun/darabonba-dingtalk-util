/**
 * This is for DingUtil
 */
// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.Security.Cryptography;
using System.Text;
using Tea.Utils;


namespace AlibabaCloud.Ding.Util
{
    public class Common 
    {

        public static string GetTimestamp()
        {
            TimeSpan ts = DateTime.UtcNow - new DateTime(1970, 1, 1, 0, 0, 0, 0);
            return Convert.ToInt64(ts.TotalMilliseconds).ToString();
        }

        public static bool HasError(Dictionary<string, object> res)
        {
            if(res.Get("error_response")!=null)
            {
                return true;
            }

            int code = res.Get("errcode").ToSafeInt(0);

            if(code == 0)
            {
                return false;
            }

            return true;
        }

        public static string ComputeSignature(string accessSecret, string canonicalString)
        {
            byte[] signData;
            using (KeyedHashAlgorithm algorithm = CryptoConfig.CreateFromName("HMACSHA256") as KeyedHashAlgorithm)
            {
                algorithm.Key = Encoding.UTF8.GetBytes(accessSecret);
                signData = algorithm.ComputeHash(Encoding.UTF8.GetBytes(canonicalString.ToCharArray()));
            }
            string signedStr = Convert.ToBase64String(signData);
            return signedStr;
        }

        public static string GetCanonicalStringForIsv(string timestamp, string suiteTicket)
        {
            string result = timestamp;
            if(suiteTicket != null)
            {
                result += "\n" + suiteTicket;
            }

            return result;
        }


    }
}
