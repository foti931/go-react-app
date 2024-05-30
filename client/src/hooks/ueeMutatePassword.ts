import { useMutation } from '@tanstack/react-query'
import { useError } from './useError'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

export const useMutatePassword = () => {
  const { switchErrorHandling } = useError()
  const navigate = useNavigate()

  const forgotPasswordMutation = useMutation(
    async (email: { email: string }) =>
      await axios.post(
        `${process.env.REACT_APP_API_URL}/password/forgot`,
        email
      ),
    {
      onSuccess() {
        navigate('/')
      },
      onError: (error: any) => {
        if (error.response.data.message) {
          switchErrorHandling(error.response.data.message)
        } else {
          switchErrorHandling(error.response.data)
        }
      },
    }
  )

  const resetPasswordMutation = useMutation(
    async (password: { password: string; token: string }) =>
      await axios.post(
        `${process.env.REACT_APP_API_URL}/password/reset`,
        password
      ),
    {
      onSuccess() {
        navigate('/')
      },
      onError: (error: any) => {
        if (error.response.data.message) {
          switchErrorHandling(error.response.data.message)
        } else {
          switchErrorHandling(error.response.data)
        }
      },
    }
  )

  return { forgotPasswordMutation, resetPasswordMutation }
}
