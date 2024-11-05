import { axiosInstance } from '@/axios'



// User Section
export async function getAllUsers() {
  try {
    const res = await axiosInstance.get('api/v1/users')
    return res
  } catch (error) {
    console.error('API request failed:', error)
    throw error
  }
}

export async function updateUser(user: any) {
  let res = await axiosInstance.patch(`api/v1/users/${user.ID}`, user)
  return res.data
}

export async function getUser(index: number) {
  let res = await axiosInstance.get(`api/v1/users/${index}`)
  if (res.status == 200) {
    console.log('Successfully Retrieved User Details')
    return res.data
  }
  }
export async function addUser(userId: number, user: any) {
  let res = await axiosInstance.post(`api/v1/users/${userId}/users`, user)
  if (res.status == 200) {
    console.log('Successfully Created a User')
    return res
  }
}
export async function getAllUserImages(index: number) {
  try {
    const res = await axiosInstance.get(`api/v1/users/${index}/images`)
    if (res.status === 200) {
      return res.data
    }
  } catch (error) {
    console.error('API request failed:', error)
    throw error
  }
}
