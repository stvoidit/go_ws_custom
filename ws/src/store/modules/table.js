export default {
    namespaced: true,
    state: {
        table: []
    },
    getters: {
        GetTable: state => state.table
    },
    mutations: {
        _onmessageTable: (state, data) => {
            state.table = data.table
        },
    },
    actions: {
        onmessageTable: (ctx, event) => {
            ctx.commit("_onmessageTable", JSON.parse(event.data))
        }
    }
}