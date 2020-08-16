<template>
    <div id="app" class="container">
        <div class="pure-u-1">
            <p>{{GetTicker}}</p>
        </div>
        <br />
        <div class="pure-u-1">
            <button class="pure-button" @click="Send">send</button>
        </div>
        <br />
        <br />
        <div class="pure-u-1">
            <table class="pure-table pure-table-bordered">
                <thead>
                    <tr>
                        <th v-for="(h, i) in headers" :key="i">{{h.text}}</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(row, i) in GetTable" :key="i">
                        <td>{{row.num}}</td>
                        <td>{{row.text}}</td>
                        <td>
                            <button @click="Del(i)">delete</button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
    name: "App",
    data() {
        return {
            wsTable: new WebSocket(`ws://${location.hostname}:9999/table`),
            wsTicker: new WebSocket(`ws://${location.hostname}:9999/timer`),
            headers: [
                { text: "321", value: "num" },
                { text: "123", value: "text" },
                { text: "delete" },
            ],
        };
    },
    created() {
        this.wsTable.onmessage = this.reciveTable;
        this.wsTicker.onmessage = this.reciveTicker;
    },
    computed: {
        ...mapGetters("table", ["GetTable"]),
        ...mapGetters("ticker", ["GetTicker"]),
    },
    methods: {
        ...mapActions("table", {
            reciveTable: "onmessage",
        }),
        ...mapActions("ticker", {
            reciveTicker: "onmessage",
        }),
        Send() {
            let msg = { text: "hello!", num: "11" };
            this.wsTable.send(JSON.stringify(msg));
        },
        Del(index) {
            console.log(index);
            this.wsTable.send(JSON.stringify({ index: index.toString() }));
        },
    },
};
</script>
<style >
.container {
    margin: 0 5%;
}
</style>


