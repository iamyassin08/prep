import { axiosInstance } from '@/axios'



// Product Section
export async function getAllProducts() {
  try {
    const res = await axiosInstance.get('api/v1/products')
    return res
  } catch (error) {
    console.error('API request failed:', error)
    throw error
  }
}

export async function updateProduct(product: any) {
  let res = await axiosInstance.patch(`api/v1/products/${product.ID}`, product)
  return res.data
}

export async function getProduct(index: number) {
  let res = await axiosInstance.get(`api/v1/products/${index}`)
  if (res.status == 200) {
    console.log('Successfully Retrieved Product Details')
    return res.data
  }
  }
export async function addProduct(categoryId: number, product: any) {
  let res = await axiosInstance.post(`api/v1/categories/${categoryId}/products`, product)
  if (res.status == 200) {
    console.log('Successfully Created a Product')
    return res
  }
}
export async function getAllProductImages(index: number) {
  try {
    const res = await axiosInstance.get(`api/v1/products/${index}/images`)
    if (res.status === 200) {
      return res.data
    }
  } catch (error) {
    console.error('API request failed:', error)
    throw error
  }
}
