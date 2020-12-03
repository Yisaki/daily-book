import request from './request'

export function PeriodDto(){
    this.id=0,
    this.type=0,
    this.happenTime=''
}

export function getLastRecord(){
    return request({
        url:'/period/getLastRecord',
        method:'GET'
    })
}

export function saveRecord(period){
    console.log(period)
    console.log(JSON.stringify(period))
    return request({
        url:'/period/saveRecord',
        method:'POST',
        data:JSON.stringify(period)
    })
}

export function listRecord(pageDto){
    return request({
        url:'/period/listRecord',
        method:'POST',
        data:JSON.stringify(pageDto)
    })
}
