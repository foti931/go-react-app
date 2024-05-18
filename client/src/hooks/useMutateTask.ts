import { useNavigate } from 'react-router-dom'
import useStore from '../store'
import { useError } from './useError'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import axios from 'axios'
import { Task } from '../types'

export const useMutateTasks = () => {
  const queryClient = useQueryClient()
  const resetEditedTask = useStore((state) => state.resetEditedTask)
  const { switchErrorHandling } = useError()

  const createTaskMutation = useMutation(
    async (task: Omit<Task, 'id' | 'created_at' | 'updated_at'>) =>
      await axios.post<Task>(`${process.env.REACT_APP_API_URL}/tasks`, task),
    {
      onSuccess: (res) => {
        const previousTasks = queryClient.getQueryData<Task[]>(['tasks'])
        if (previousTasks) {
          queryClient.setQueryData<Task[]>(
            ['tasks'],
            [...previousTasks, res.data]
          )
          resetEditedTask()
        }
      },
      onError: (error: any) => {
        if (error.response.data.message) {
          console.log(error)
          switchErrorHandling(error.response.data.message)
        } else {
          console.log(error)
          switchErrorHandling(error.response.data)
        }
      },
    }
  )

  const updateTaskMutation = useMutation(
    async (task: Omit<Task, 'created_at' | 'updated_at'>) =>
      await axios.put<Task>(
        `${process.env.REACT_APP_API_URL}/tasks/${task.id}`,
        { id: task.id, title: task.title }
      ),
    {
      onSuccess: (res) => {
        const previousTasks = queryClient.getQueryData<Task[]>(['tasks'])
        if (previousTasks) {
          console.log(res.data)
          queryClient.setQueryData<Task[]>(
            ['tasks'],
            previousTasks.map((task) =>
              task.id === res.data.id ? res.data : task
            )
          )
          resetEditedTask()
        }
      },
      onError: (error: any) => {
        if (error.response.data.message) {
          console.log(error)
          switchErrorHandling(error.response.data.message)
        } else {
          console.log(error)
          switchErrorHandling(error.response.data)
        }
      },
    }
  )

  const deleteTaskMutation = useMutation(
    async (id: number) =>
      await axios.delete(`${process.env.REACT_APP_API_URL}/tasks/${id}`),
    {
      onSuccess: (_, id) => {
        const previousTasks = queryClient.getQueryData<Task[]>(['tasks'])
        if (previousTasks) {
          queryClient.setQueryData<Task[]>(
            ['tasks'],
            previousTasks.filter((task) => task.id !== id)
          )
        }
      },
      onError: (error: any) => {
        if (error.response.data.message) {
          console.log(error)
          switchErrorHandling(error.response.data.message)
        } else {
          console.log(error)
          switchErrorHandling(error.response.data)
        }
      },
    }
  )

  return { createTaskMutation, updateTaskMutation, deleteTaskMutation }
}
