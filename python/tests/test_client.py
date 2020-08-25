import unittest
from alibabacloud_ding_util.client import Client


class TestClient(unittest.TestCase):
    def test_get_timestamp(self):
        self.assertEqual(13, len(Client.get_timestamp()))

    def test_has_error(self):
        he = Client.has_error
        res = {
            'error_response': '',
            'errcode': None
        }
        self.assertEqual(True, he(res))

        res['error_response'] = None
        self.assertEqual(False, he(res))

        res['errcode'] = '1'
        self.assertEqual(True, he(res))

        res['errcode'] = 'test'
        self.assertEqual(False, he(res))

        res['errcode'] = '0'
        self.assertEqual(False, he(res))

    def test_compute_signature(self):
        self.assertEqual(
            'Aymga2LNFrM+tnkr6MYLFY2Jou46h2/Omogeu0iMCRQ=',
            Client.compute_signature('secret', 'test')
        )

    def test_get_canonical_string_for_isv(self):
        res = Client.get_canonical_string_for_isv('10000', 'suite_ticket')
        self.assertEqual('\nsuite_ticket', res)

        res = Client.get_canonical_string_for_isv('10000', None)
        self.assertEqual('10000', res)