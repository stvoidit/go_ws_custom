export default {
    namespaced: true,
    state: {
        tick: null
    },
    getters: {
        GetTicker: state => Date(state.tick)
    },
    mutations: {
        _onmessage: (state, data) => {
            state.tick = data.tick
        }
    },
    actions: {
        onmessage: (ctx, event) => {
            ctx.commit("_onmessage", JSON.parse(event.data))
        }
    }
}