import time
import hashlib
import base64
import hmac


class Client:
    @staticmethod
    def get_timestamp():
        return str(int(round(time.time() * 1000)))

    @staticmethod
    def has_error(res):
        if res.get('error_response') is not None:
            return True

        if res.get('errcode') is None:
            return False

        try:
            n = int(res.get('errcode'))
            if n == 0:
                return False
        except ValueError:
            return False
        else:
            return True

    @staticmethod
    def compute_signature(secret, string):
        hash_val = hmac.new(secret.encode('utf-8'), string.encode('utf-8'), hashlib.sha256).digest()
        signature = base64.b64encode(hash_val).decode('utf-8')
        return signature

    @staticmethod
    def get_canonical_string_for_isv(timestamp, suite_ticket):
        if suite_ticket is not None:
            return '\n%s' % suite_ticket
        else:
            return timestamp
