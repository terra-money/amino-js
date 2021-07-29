import * as Amino from "../";

const txData =
    "1QPGwQI/Cs0CaHynGgqwApt/+gkKLVByb3Bvc2FsOiBQYXlXaXRoVGVycmEgQ29tbXVuaXR5IFBvb2wgRnVuZGluZxLPAVdlIG5lZWQgc29tZSBmdW5kcyB0byByZXZpc2lvbiBvZiB0aGUgc3lzdGVtIGFuZCBwbHVnaW5zIChjb2RpbmcpLCBkZXNpZ24gYW5kIGlsbHVzdHJhdGlvbnMsIGxlZ2FsLiBEZXRhaWxlZCBkZXNjcmlwdGlvbiBvbiB0b3BpYzoKaHR0cHM6Ly9hZ29yYS50ZXJyYS5tb25leS90L3Byb3Bvc2FsLXBheXdpdGh0ZXJyYS1jb21tdW5pdHktcG9vbC1mdW5kaW5nLzQ5MhoUfq76xy5LXmHN3Oisl9AeVhO/WdIiEwoFdWx1bmESCjIzMDAwMDAwMDAaFH6u+scuS15hzdzorJfQHlYTv1nSEhMKDQoFdWx1bmESBDI5MTQQqdkPGmoKJuta6YchA2322oIb/PDKei/IOOqGNmlNQifQqIwkItGd4gmzoWf/EkDQuLnre9spJJ42gmzEJvbdcs3PEV02Um1kfIfy7HUXMmAH9WyNV4MZjtMwMDJE27kWV8prJnUGTuMux+ryIAnx";

const tx = {
    type: "core/StdTx",
    value: {
        msg: [
            {
                type: "gov/MsgSubmitProposal",
                value: {
                    content: {
                        type: "distribution/CommunityPoolSpendProposal",
                        value: {
                            title: "Proposal: PayWithTerra Community Pool Funding",
                            description:
                                "We need some funds to revision of the system and plugins (coding), design and illustrations, legal. Detailed description on topic:\nhttps://agora.terra.money/t/proposal-paywithterra-community-pool-funding/492",
                            recipient: "terra106h043ewfd0xrnwuazkf05q72cfm7kwj784la2",
                            amount: [{ denom: "uluna", amount: "2300000000" }],
                        },
                    },
                    initial_deposit: null,
                    proposer: "terra106h043ewfd0xrnwuazkf05q72cfm7kwj784la2",
                },
            },
        ],
        fee: { amount: [{ denom: "uluna", amount: "2914" }], gas: "257193" },
        signatures: [
            {
                pub_key: {
                    type: "tendermint/PubKeySecp256k1",
                    value: "A2322oIb/PDKei/IOOqGNmlNQifQqIwkItGd4gmzoWf/",
                },
                signature:
                    "0Li563vbKSSeNoJsxCb23XLNzxFdNlJtZHyH8ux1FzJgB/VsjVeDGY7TMDAyRNu5FlfKayZ1Bk7jLsfq8iAJ8Q==",
            },
        ],
        memo: "",
    },
};

describe("Tx3", () => {
    describe("decoding", () => {
        describe("Tx", () => {
            it("decodes bytes", () => {
                const bytes = Amino.base64ToBytes(txData);
                const value = Amino.unmarshalTx(bytes, true);
                console.log(value);
                expect(value).toMatchObject(tx);
            });
        });
    });

    describe("encoding", () => {
        describe("Tx", () => {
            it("encodes value", () => {
                const bytes = Amino.marshalTx(tx, true);
                const data = Amino.bytesToBase64(bytes);
                expect(data).toBe(txData);
            });
        });
    });
});
