import { axiosInstance } from '@/axios'


// Get all users
export async function getAllUsers() {
  try {
    const res = await axiosInstance.get('api/v1/users')
    return res.data
  } catch (error) {
    console.error('API request failed:', error)
    throw error
  }
}

// Get user by ID
export async function getUser(userId: number) {
  try {
    const res = await axiosInstance.get(`api/v1/users/${userId}`)
    if (res.status === 200) {
      console.log('Successfully Retrieved User Details')
      return res.data
    }
  } catch (error) {
    console.error('API request failed:', error)
    throw error
  }
}

// Update user by ID
export async function updateUser(user: any) {
  try {
    const res = await axiosInstance.patch(`api/v1/users/${user.id}`, user)
    return res.data
  } catch (error) {
    console.error('API request failed:', error)
    throw error
  }
}

// Add a new user 
export async function addUser(user: any) {
  try {
    const res = await axiosInstance.post('api/v1/users', user) 
    if (res.status === 201) {
      console.log('Successfully Created a User')
      return res.data
    }
  } catch (error) {
    console.error('API request failed:', error)
    throw error
  }
}

// Delete a new user 
export async function DeleteUser(user: any) {
  try {
    const res = await axiosInstance.delete('api/v1/users', user) 
    if (res.status === 201) {
      console.log('Successfully Deleted a User')
      return res.data
    }
  } catch (error) {
    console.error('API request failed:', error)
    throw error
  }
}
