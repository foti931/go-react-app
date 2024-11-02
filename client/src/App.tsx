import { useEffect } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Auth } from './components/Auth';
import { ToDo } from './components/ToDo';
import { PasswordForgot } from './components/password/Forgot';
import axios from 'axios';
import { CsrfToken } from './types';
import 'semantic-ui-css/semantic.min.css'
import { PasswordReset } from './components/password/Reset';
import { NavBar } from "./components/NavBar";

function App() {

  useEffect(() => {
    axios.defaults.withCredentials = true;
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(`${process.env.REACT_APP_API_URL}/csrf`);
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token;
    }
    getCsrfToken();
  }, [])

  return (
      <>
        <div>
          <NavBar />
          <BrowserRouter>
            <Routes>
              <Route path="/" element={<Auth />} />
              <Route path="/todo" element={<ToDo />} />
              <Route path="/password/forgot" element={<PasswordForgot />} />
              <Route path="/password/reset" element={<PasswordReset />} />
              <Route path="/password" element={<PasswordForgot />} />
            </Routes>
          </BrowserRouter>
         </div>
       </>
  );
}

export default App;