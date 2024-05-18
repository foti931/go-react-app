import axios from "axios"
import { CsrfToken } from "../types"
import useStore from "../store"
import { useNavigate } from "react-router-dom"

export const useError = () => {
    const navigate = useNavigate()
    const resetEditedTask = useStore((state) => state.resetEditedTask)
    const getCsrfToken = async () => {
        try {
            const { data } = await axios.get<CsrfToken>(`${process.env.REACT_APP_API_URL}/csrf`)
            axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
        } catch (error: any) {
            navigate('/')
            resetEditedTask()
        }
    }

    const switchErrorHandling = (msg: string): void => {
        switch (msg) {
            case 'invalid csrf token':
                getCsrfToken()
                alert('Invalid CSRF token. Please reload the page.')
                break;
            case 'invalid or expired jwt':
                alert('Invalid or expired JWT. Please login again.')
                resetEditedTask()
                navigate('/')
                break
            case 'missing or malformed jwt':
                alert('Missing or malformed JWT. Please login again.')
                resetEditedTask()
                navigate('/')
                break
            case msg.match('duplicate key value violates') && msg:
                alert('email is already taken, please use another email')
                navigate('/')
                break
            case 'crypto/bcrypt: hashedPassword is not the hash of the given password':
                alert('emmail or password is incorrect')
                break
            case 'record not found':
                alert('emmail or password is incorrect')
                break
            default:
                alert(msg)
                break;
        }
    }

    return { getCsrfToken, switchErrorHandling }
}