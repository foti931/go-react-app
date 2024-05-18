import axios from 'axios'
import { useError } from './useError'
import { Task } from '../types'
import { useQuery } from '@tanstack/react-query'

export const useQueryTasks = () => {
  const { switchErrorHandling } = useError()

  const getTasks = async () => {
    const { data } = await axios.get<Task[]>(
      `${process.env.REACT_APP_API_URL}/tasks`,
      { withCredentials: true }
    )
    return data
  }

  return useQuery<Task[], Error>({
    queryKey: ['tasks'],
    queryFn: getTasks,
    staleTime: Infinity,
    onError: (error: any) => {
      if (error.response.data.message) {
        switchErrorHandling(error.response.data.message)
      } else {
        switchErrorHandling(error.response.data)
      }
    },
  })
}
