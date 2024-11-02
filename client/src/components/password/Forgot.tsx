import { FormEvent, useState } from "react"
import { useMutatePassword } from "../../hooks/ueeMutatePassword"

export const PasswordForgot = () => {
    const [email, setEmail] = useState('')
    const { forgotPasswordMutation } = useMutatePassword()

    const submitPasswordHandler = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        await forgotPasswordMutation.mutateAsync({ email })
            .then(() => {
                alert('成功')
            })
            .catch(() => {
                alert('Failed to send an email')
            })
    }

    return (
        <div className="flex justify-center items-center flex-col min-h-screen text-grey-600 font-mono">
            <div className="my-5 flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                    <path fillRule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12Zm11.378-3.917c-.89-.777-2.366-.777-3.255 0a.75.75 0 0 1-.988-1.129c1.454-1.272 3.776-1.272 5.23 0 1.513 1.324 1.513 3.518 0 4.842a3.75 3.75 0 0 1-.837.552c-.676.328-1.028.774-1.028 1.152v.75a.75.75 0 0 1-1.5 0v-.75c0-1.279 1.06-2.107 1.875-2.502.182-.088.351-.199.503-.331.83-.727.83-1.857 0-2.584ZM12 18a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5Z" clipRule="evenodd" />
                </svg>
                <span className="font-bold text-2xl">パスワードを忘れた</span>
            </div>
            <form onSubmit={submitPasswordHandler}>
                <div>
                    <p>登録済みのメールアドレスを入力してください。</p>
                    <input type="email" className="px-4 py-2 w-full border rounded-md" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
                </div>
                <div className="flex">
                    <a href="/" className="w-full py-2 mt-4 mx-2 bg-slate-300 hover:bg-stale-600 rounded-md  text-white  text-center">戻る</a>
                    <button className="w-full py-2 mt-4 bg-blue-600 hover:bg-green-700 rounded-md text-white disabled:bg-blue-100" disabled={!email} type="submit">メール送信</button>
                </div>
            </form>
        </div >
    )
}