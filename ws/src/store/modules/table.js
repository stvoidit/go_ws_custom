export default {
    namespaced: true,
    state: {
        table: []
    },
    getters: {
        GetTable: state => state.table
    },
    mutations: {
        _onmessage: (state, data) => {
            state.table = data.table
        },
    },
    actions: {
        onmessage: (ctx, event) => {
            ctx.commit("_onmessage", JSON.parse(event.data))
        }
    }
}