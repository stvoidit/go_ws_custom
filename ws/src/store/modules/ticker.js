export default {
    namespaced: true,
    state: {
        tick: null
    },
    getters: {
        GetTicker: state => Date(state.tick)
    },
    mutations: {
        _onmessageTicker: (state, data) => {
            state.tick = data.tick
        }
    },
    actions: {
        onmessageTicker: (ctx, event) => {
            ctx.commit("_onmessageTicker", JSON.parse(event.data))
        }
    }
}