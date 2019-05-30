import axios from 'axios'

const baseURL = 'http://192.168.1.101:3002/web'


const api = {
    install(Vue){
        Vue.prototype.$api = axios.create({
            baseURL,
            transformResponse: function (data) {
                data = JSON.parse(data)
                if(data.code === 401){
                    console.log('llll')
                    axios.get(`${baseURL}?path=${window.router.history.current.path.split('/')[window.router.history.current.path.split('/').length-1]}`).then((res)=>{
                        window.location = res.data.result
                    })
                }
                return data
            }
        })
        Vue.prototype.$setSessionToken = (sessionToken) => {
            Vue.prototype.$api = axios.create({
                baseURL,
                headers:{
                    'token':sessionToken
                },
                transformResponse: function (data) {
                    data = JSON.parse(data)
                    if(data.code === 401){
                        axios.get(`${baseURL}?path=${window.router.history.current.path.split('/')[window.router.history.current.path.split('/').length-1]}`).then((res)=>{
                            window.location = res.data.result
                        })
                    }
                    return data
                }
            })
        }
    }
}

export default api