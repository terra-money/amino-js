import * as Amino from '../';

const txData = 'qgHwYl3uCjIrTd2SChRXmZp0ijNn4Ck45U3Dn8935ucetBIQCgV1bHVuYRIHMTAwMDAwMBoEdWtydxIEEMCEPRpqCibrWumHIQMyy6qxoDl1ZfPvJtzvbYmXSm+PoEMy5aBazEVZkWwbFxJAnhH429iL4VD3K4nADVf8sjD7LdPzuuM53QrBh98sENIs8NUpPJNTUQ9zxQFidWEaFFBjP12x74fpOxsypwmW5w==';

const tx = {
    'type':  'auth/StdTx',
    'value': {
        'msg':        [{
            'type':  'market/MsgSwap',
            'value': {
                'trader': 'cosmos127ve5ay2xdn7q2fcu4xu8870wlnww845pv6ten',
                'offer_coin':  {
                    'denom': 'uluna',
                    'amount': '1000000',
                },
                'ask_denom':   'ukrw',
            }
        }],
        'fee':        {
            'amount': null,
            'gas':    '1000000'
        },
        'signatures': [
            {
                'signature': 'nhH429iL4VD3K4nADVf8sjD7LdPzuuM53QrBh98sENIs8NUpPJNTUQ9zxQFidWEaFFBjP12x74fpOxsypwmW5w==',
                'pub_key':   {
                    'type':  'tendermint/PubKeySecp256k1',
                    'value': 'AzLLqrGgOXVl8+8m3O9tiZdKb4+gQzLloFrMRVmRbBsX'
                }
            }
        ],
        'memo':       ''
    }
};

describe('Extensions', () => {
    describe('decoding', () => {
        describe('Tx', () => {
            it('decodes bytes', () => {
                const bytes = Amino.base64ToBytes(txData);
                const value = Amino.unmarshalTx(bytes, true);
                expect(value).toMatchObject(tx);
            });
        });
    });

    describe('encoding', () => {
        describe('Tx', () => {
            it('encodes value', () => {
                const bytes = Amino.marshalTx(tx, true);
                const data  = Amino.bytesToBase64(bytes);
                expect(data).toBe(txData);
            });
        });
    });
});
