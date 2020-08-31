/**
 * This is for DingUtil
 */
import * as $tea from '@alicloud/tea-typescript';
import sha256 from 'crypto-js/hmac-sha256';
import encBase64 from 'crypto-js/enc-base64'

export default class Client {

  static getTimestamp(): string {
    return `${(new Date()).getTime()}`;
  }

  static hasError(res: {[key: string ]: any}): boolean {
    if (!res) {
      return true;
    }
    if (res.error_response) {
      return true;
    }
    if (!res.errcode) {
      return false
    }

    return true;
  }

  static computeSignature(accessSecret: string, canonicalString: string): string {
    const hash = sha256(canonicalString, accessSecret);
    return encBase64(hash.toString());
  }

  static getCanonicalStringForIsv(timestamp: string, suiteTicket: string): string {
    let res = timestamp
    if (suiteTicket) {
      res = `${res}\n${suiteTicket}`
    }
    
    return res;
  }

}
