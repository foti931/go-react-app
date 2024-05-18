import { FormEvent, useState } from "react"
import { useMutateAuth } from "../hooks/useMutateAuth"
import { ArrowPathIcon, CheckBadgeIcon } from "@heroicons/react/24/solid"

export const Auth = () => {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [isLogin, setIsLogin] = useState(true)
    const { loginMutating, registerMutating } = useMutateAuth()

    const submitAuthHandler = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (isLogin) {
            loginMutating.mutate({ email, password })
        } else {
            await registerMutating
                .mutateAsync({ email, password })
                .then(() => {
                    loginMutating.mutate({ email, password })
                }).catch((error: any) => {
                    alert('Failed to create a new account')
                })
        }
    }
    return (
        <div className="flex justify-center items-center flex-col min-h-screen text-grey-600 font-mono">
            <div className="flex items-center">
                <CheckBadgeIcon className="w-8 h-8 mr-2 text-blue-500" />
                <span className="font-bold text-2xl">TODO App</span>
            </div>
            <h2 className="my-6">{isLogin ? 'Login' : 'Create a new account'}</h2>
            <form onSubmit={submitAuthHandler}>
                <div>
                    <input type="email" className="px-4 py-2 border rounded-md" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
                </div>
                <div>
                    <input type="password" className="px-4 py-2 border rounded-md" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
                </div>
                <div>
                    <button className="w-full py-2 mt-4 bg-blue-600 hover:bg-green-700 rounded-md text-white" disabled={!email || !password} type="submit">{isLogin ? 'Login' : 'Create Account'}</button>
                </div>
            </form>
            <div className="flex items-center">
                <ArrowPathIcon className="w-6 h-6 my-2 text-blue-600 cursor-pointer" onClick={() => setIsLogin(!isLogin)} />
                <span className="text-sm">{`Swith to ${isLogin ? "crating a new account" : "Login"}`}</span>
            </div>
        </div>
    )
}
