<template>
    <div>
        <div class="pure-u-1">
            <button class="pure-button" @click="Add">add</button>
        </div>
        <table class="pure-table pure-table-bordered">
            <thead>
                <tr>
                    <th v-for="(field, i) in headers" :key="i">{{field}}</th>
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
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
    name: "Table",
    data() {
        return {
            ws: new WebSocket(`ws://${location.hostname}:9999/table`),
            headers: ["number", "text", "delete"],
        };
    },
    created() {
        this.ws.onmessage = this.receive;
    },
    computed: {
        ...mapGetters("table", ["GetTable"]),
    },
    methods: {
        ...mapActions("table", {
            receive: "onmessage",
        }),
        Add() {
            let msg = { text: "hello!", num: "11" };
            this.ws.send(JSON.stringify(msg));
        },
        Del(index) {
            this.ws.send(JSON.stringify({ index: index.toString() }));
        },
    },
};
</script>
