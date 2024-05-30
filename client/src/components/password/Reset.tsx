import { useState } from "react"

export const PasswordReset = () => {

    const [password, setPassword] = useState('')
    const [confirmPassword, setConfirmPassword] = useState('')

    const submitPasswordResetHandler = (e: React.FormEvent) => {
        e.preventDefault()
    }

    return (
        <div className="flex justify-center items-center flex-col min-h-screen text-grey-600 font-mono">
            <div className="my-5 flex items-center">

                <span className="font-bold text-2xl">パスワードのリセットを行います。</span>
            </div>
            <form onSubmit={submitPasswordResetHandler}>
                <div>
                    <p className="text-slate-400 m-0">新しいパスワード</p>
                    <input type="password" className="px-4 py-2 mb-6 w-full border rounded-md" placeholder="パスワード" value={password} onChange={(e) => setPassword(e.target.value)} />
                    <p className="text-slate-400 m-0">確認用</p>
                    <input type="password" className="px-4 py-2 w-full border rounded-md" placeholder="パスワード確認" value={confirmPassword} onChange={(e) => setConfirmPassword(e.target.value)} />
                </div>
                <div className="flex items-center">
                    <button className="w-full py-2 mt-4 bg-blue-600 hover:bg-green-700 rounded-md text-white disabled:bg-blue-100" disabled={!password || !confirmPassword} type="submit">パスワードリセット</button>
                </div>
            </form>
        </div >
    )
}
