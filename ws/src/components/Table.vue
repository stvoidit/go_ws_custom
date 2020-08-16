<template>
    <div>
        <div class="pure-u-1">
            <button class="pure-button" @click="Add">add</button>
        </div>
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
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
    name: "Table",
    data() {
        return {
            wsTable: new WebSocket(`ws://${location.hostname}:9999/table`),
            headers: [
                { text: "321", value: "num" },
                { text: "123", value: "text" },
                { text: "delete" },
            ],
        };
    },
    created() {
        this.wsTable.onmessage = this.reciveTable;
    },
    computed: {
        ...mapGetters("table", ["GetTable"]),
    },
    methods: {
        ...mapActions("table", {
            reciveTable: "onmessageTable",
        }),
        Add() {
            let msg = { text: "hello!", num: "11" };
            this.wsTable.send(JSON.stringify(msg));
        },
        Del(index) {
            this.wsTable.send(JSON.stringify({ index: index.toString() }));
        },
    },
};
</script>
