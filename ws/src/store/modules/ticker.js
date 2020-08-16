export default {
    namespaced: true,
    state: {
        tick: new Date()
    },
    getters: {
        GetTicker: state => state.tick
    },
    mutations: {
        _onmessageTicker: (state, data) => {
            state.tick = new Date(Date.parse(data.tick))
        }
    },
    actions: {
        onmessage: (ctx, event) => {
            ctx.commit("_onmessageTicker", JSON.parse(event.data))
        }
    }
}