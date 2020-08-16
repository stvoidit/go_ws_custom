<template>
    <div class="pure-u-1">
        <p>{{GetTicker}}</p>
    </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
    name: "Ticker",
    data() {
        return {
            wsTicker: new WebSocket(`ws://${location.hostname}:9999/timer`),
        };
    },
    created() {
        this.wsTicker.onmessage = this.reciveTicker;
    },
    computed: {
        ...mapGetters("ticker", ["GetTicker"]),
    },
    methods: {
        ...mapActions("ticker", {
            reciveTicker: "onmessageTicker",
        }),
    },
};
</script>
