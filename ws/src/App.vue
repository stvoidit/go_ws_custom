<template>
    <div id="app" class="container">
        <Ticker />
        <br />
        <div class="pure-u-1">
            <Table />
        </div>
    </div>
</template>

<script>
import Table from "@/components/Table";
import Ticker from "@/components/Ticker";
export default {
    name: "App",
    components: { Table, Ticker },
};
</script>
<style >
.container {
    margin: 0 5%;
}
</style>


