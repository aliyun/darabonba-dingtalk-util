using System;
using System.Collections.Generic;
using AlibabaCloud.Ding.Util;
using Xunit;

namespace tests
{
    public class CommonTest
    {
        [Fact]
        public void TestGetTimestamp()
        {
            Assert.Equal(13, Common.GetTimestamp().Length);
        }

        [Fact]
        public void TestHasError()
        {
            Dictionary<string, object> dict = new Dictionary<string, object>();
            dict.Add("error_response", "");
            dict.Add("errcode", null);
            Assert.True(Common.HasError(dict));

            dict["error_response"] = null;
            Assert.False(Common.HasError(dict));

            dict["errcode"] = "1";
            Assert.True(Common.HasError(dict));

            dict["errcode"] = "test";
            Assert.False(Common.HasError(dict));

            dict["errcode"] = "0";
            Assert.False(Common.HasError(dict));
        }

        [Fact]
        public void TestComputeSignature()
        {
            Assert.Equal("Aymga2LNFrM+tnkr6MYLFY2Jou46h2/Omogeu0iMCRQ=", Common.ComputeSignature("secret", "test"));
        }

        [Fact]
        public void TestGetCanonicalStringForIsv()
        {
            Assert.Equal("\nsuite_ticket", Common.GetCanonicalStringForIsv("10000", "suite_ticket"));

            Assert.Equal("10000", Common.GetCanonicalStringForIsv("10000", null));
        }
    }
}
