import service from '@/utils/request'

export const createDayzItem = (data) => {
  return service({
    url: '/dayzItem/createDayzItem',
    method: 'post',
    data,
  })
}

export const updateDayzItem = (data) => {
  return service({
    url: '/dayzItem/updateDayzItem',
    method: 'put',
    data,
  })
}

export const deleteDayzItemByIds = (data) => {
  return service({
    url: '/dayzItem/deleteDayzItemByIds',
    method: 'delete',
    data,
  })
}

export const getDayzItem = (data) => {
  return service({
    url: '/dayzItem/getDayzItem',
    method: 'post',
    data,
  })
}

export const getDayzItemByID = (data) => {
  return service({
    url: '/dayzItem/getDayzItemByID',
    method: 'get',
    data,
  })
}

export const getDayzItemList = (data) => {
  return service({
    url: '/dayzItem/getDayzItemList',
    method: 'post',
    data,
  })
}
