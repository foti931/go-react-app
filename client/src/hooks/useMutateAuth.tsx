import { useNavigate } from "react-router-dom";
import useStore from "../store";
import { useError } from "./useError";
import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import { Credential } from "../types";

export const useMutateAuth = () => {
    const navigate = useNavigate()
    const resetEditedTask = useStore((state) => state.resetEditedTask)
    const { switchErrorHandling } = useError()

    const loginMutating = useMutation(
        async (user: Credential) =>
            await axios.post(`${process.env.REACT_APP_API_URL}/login`, user),
        {
            onSuccess: () => {
                navigate('/todo');
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

    const registerMutating = useMutation(
        async (user: Credential) =>
            await axios.post(`${process.env.REACT_APP_API_URL}/signup`, user),
        {
            onError: (error: any) => {
                if (error.response.data?.message) {
                    console.log("errorです" + error)
                    switchErrorHandling(error.response.data.message)
                } else {
                    console.log(error)
                    switchErrorHandling(error.response.data)
                }
            },
        }
    )

    const logoutMutating = useMutation(
        async () =>
            await axios.post(`${process.env.REACT_APP_API_URL}/logout`),
        {
            onSuccess: () => {
                resetEditedTask()
                navigate('/')
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

    return { loginMutating, registerMutating, logoutMutating }
}