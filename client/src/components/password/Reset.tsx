import {FormEvent, useState} from "react"
import { useMutatePassword } from "../../hooks/ueeMutatePassword"
import { useNavigate, useSearchParams} from "react-router-dom";

export const PasswordReset = () => {

    const [password, setPassword] = useState('')
    const [confirmPassword, setConfirmPassword] = useState('')
    const {resetPasswordMutation} = useMutatePassword()
    const [searchParams] = useSearchParams();
    const token = searchParams.get('token') ?? '';
    const navigate = useNavigate()

    const submitPasswordResetHandler = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        await resetPasswordMutation.mutateAsync({password, confirmPassword,  token})
            .then(() => {
                navigate('/')
            })
            .catch(() => {
                alert('Failed to reset password')
            })
    }

    return (
        <div className="flex justify-center items-center flex-col min-h-screen text-grey-600 font-mono">
            <div className="my-5 flex items-center">
                <span className="font-bold text-2xl">パスワードのリセットを行います。</span>
            </div>
            <form onSubmit={submitPasswordResetHandler}>
                <div>
                    <p className="text-slate-400 m-0">新しいパスワード</p>
                    <input name="password" type="password" className="px-4 py-2 mb-6 w-full border rounded-md" placeholder="パスワード" value={password} onChange={(e) => setPassword(e.target.value)} />
                    <p className="text-slate-400 m-0">確認用</p>
                    <input name="confirm_password"　type="password" className="px-4 py-2 w-full border rounded-md" placeholder="パスワード確認" value={confirmPassword} onChange={(e) => setConfirmPassword(e.target.value)} />
                </div>
                <div className="flex items-center">
                    <button className="w-full py-2 mt-4 bg-blue-600 hover:bg-green-700 rounded-md text-white disabled:bg-blue-100" disabled={!password || !confirmPassword} type="submit">パスワードリセット</button>
                </div>
            </form>
        </div >
    )
}
