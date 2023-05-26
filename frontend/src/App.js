import './App.css';
import {Navigate} from 'react-router-dom'
import {
    GuardConfigProvider,
    GuardedRoute,
    GuardedRoutes,
} from 'react-router-guarded-routes'
import Login from './login/components/Login'
import SignUp from './login/components/SignUp'
import {useEffect} from 'react'
import LandingPage from './cv-maker/components/LandingPage'
import CVCreate from './cv-maker/components/CVCreate'
import {AuthContext} from './login/context'
import AllUserGenerated from './cv-maker/components/AllUserGenerated'

const ProtectedRoute = ({user, userLoading, children})=> {
    useEffect(() => {
        if (!userLoading && (!user || user.id === 0)) {
            window.location.href = '/login'
        }
    }, [user, userLoading])

    return children
}

const ProtectedRouteWithConsumer = ({children}) => {
    return <AuthContext.Consumer>{data => <ProtectedRoute {...data}>{children}</ProtectedRoute>}</AuthContext.Consumer>
}

const App = () =>   {
  return (
      <>
        <GuardConfigProvider>
            <GuardedRoutes>
                <GuardedRoute path='/' element={<Navigate to='/login'/>}/>
                <GuardedRoute path='login' element={<Login/>}/>
                <GuardedRoute path='sign-up' element={<SignUp/>}/>
                <GuardedRoute path='landing-page' element={<ProtectedRouteWithConsumer><LandingPage/></ProtectedRouteWithConsumer>}/>
                <GuardedRoute path='create-cv' element={<ProtectedRouteWithConsumer><CVCreate/></ProtectedRouteWithConsumer>}/>
                <GuardedRoute path='all' element={<ProtectedRouteWithConsumer><AllUserGenerated/></ProtectedRouteWithConsumer>}/>
                <GuardedRoute path='*' element={<Navigate to='/login'/>}/>
            </GuardedRoutes>
        </GuardConfigProvider>
      </>
  );
}

export default App;
