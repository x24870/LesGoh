module owner::weighted_probability {

    use std::signer;
    use std::string::{Self};
    use aptos_token::token::{TokenId};
    use owner::bone;
    use owner::pseudorandom::{rand_u64_range};

    friend owner::urn_to_earn;

    // error codes
    const EINVALID_MINT_TYPE: u64 = 1;

    // points of each part
    const ARM_1_P:     u8 = 50;
    const ARM_3_P:     u8 = 50;
    const LEG_6_P:     u8 = 50;
    const LEG_7_P:     u8 = 50;
    const HIP_13_P:    u8 = 50;
    const HIP_15_P:    u8 = 50;
    const CHEST_18_P:  u8 = 50;
    const SKULL_21_P:  u8 = 50;
    const SKULL_26_P:  u8 = 50;
    
    // name: object name, value: weight
    const ARM_1_W:      u64 = 5;
    const ARM_3_W:      u64 = 100;
    const LEG_6_W:      u64 = 50;
    const LEG_7_W:      u64 = 60;
    const HIP_13_W:     u64 = 20;
    const HIP_15_W:     u64 = 40;
    const CHEST_18_W:   u64 = 55;
    const SKULL_21_W:   u64 = 10;
    const SKULL_26_W:   u64 = 25;
    const SUM_OF_W:     u64 = 365;

    // accumulate the weight
    struct AccumulateWeight has store, key {
        arm_1:      u64,
        arm_3:      u64,
        leg_6:      u64,
        leg_7:      u64,
        hip_13:     u64,
        hip_15:     u64,
        chest_18:   u64,
        skull_21:   u64,
        skull_26:   u64,
    }

    public(friend) fun init_weighted_probability(sign: &signer) {
        init_accumulate_weight(sign);
    }

    fun init_accumulate_weight(sender: &signer) {
        // Don't run setup more than once
        if (exists<AccumulateWeight>(signer::address_of(sender))) {
            return
        };

        let aw = AccumulateWeight {
            arm_1: 0,
            arm_3: 0,
            leg_6: 0,
            leg_7: 0,
            hip_13: 0,
            hip_15: 0,
            chest_18: 0,
            skull_21: 0,
            skull_26: 0,
        };

        aw.arm_1 = ARM_1_W;
        aw.arm_3 = aw.arm_1 + ARM_3_W;
        aw.leg_6 = aw.arm_3 + LEG_6_W;
        aw.leg_7 = aw.leg_6 + LEG_7_W;
        aw.hip_13 = aw.leg_7 + HIP_13_W;
        aw.hip_15 = aw.hip_13 + HIP_15_W;
        aw.chest_18 = aw.hip_15 + CHEST_18_W;
        aw.skull_21 = aw.chest_18 + SKULL_21_W;
        aw.skull_26 = aw.skull_21 + SKULL_26_W;

        move_to(sender, aw);
    }

    
    public(friend) fun mint_by_weight(
        sign: &signer,
        resource: &signer,
    ): (TokenId, u64) acquires AccumulateWeight {
        let aw = borrow_global<AccumulateWeight>(@owner);
        let rand_num = rand_u64_range(&signer::address_of(sign), 0, SUM_OF_W);

        let token_id: TokenId;
        let amount: u64 = 1;

        if (rand_num < aw.arm_1) {
            token_id = bone::mint(sign, resource, ARM_1_P, string::utf8(b"arm"));
        } else if (rand_num < aw.arm_3) {
            token_id = bone::mint(sign, resource, ARM_3_P, string::utf8(b"arm"));
        } else if (rand_num < aw.leg_6) {
            token_id = bone::mint(sign, resource, LEG_6_P, string::utf8(b"leg"));
        } else if (rand_num < aw.leg_7) {
            token_id = bone::mint(sign, resource, LEG_7_P, string::utf8(b"leg"));
        } else if (rand_num < aw.hip_13) {
            token_id = bone::mint(sign, resource, HIP_13_P, string::utf8(b"hip"));
        } else if (rand_num < aw.hip_15) {
            token_id = bone::mint(sign, resource, HIP_15_P, string::utf8(b"hip"));
        } else if (rand_num < aw.chest_18) {
            token_id = bone::mint(sign, resource, CHEST_18_P, string::utf8(b"chest"));
        } else if (rand_num < aw.skull_21) {
            token_id = bone::mint(sign, resource, SKULL_21_P, string::utf8(b"skull"));
        } else if (rand_num < aw.skull_26) {
            token_id = bone::mint(sign, resource, SKULL_26_P, string::utf8(b"skull"));
        } else {
            abort EINVALID_MINT_TYPE
        };

        return (token_id, amount)
    }
}