import GoogleLogin, {
  GoogleLoginResponse,
  GoogleLoginResponseOffline,
} from 'react-google-login';

function LoginGoogle() {
  const onSuccess = async (
    e: GoogleLoginResponse | GoogleLoginResponseOffline
  ) => {
    console.log(e);

    const req = await fetch('/api/auth/google');
    const data = req.json();

    console.log(data);
  };

  const onError = () => {};

  return (
    <GoogleLogin
      clientId="927290436909-tgnr0u2d86ehka37n02ivjh14s7nnhaa.apps.googleusercontent.com"
      buttonText="login"
      onSuccess={onSuccess}
      onFailure={onError}
      cookiePolicy={'single_host_origin'}
      scope="profile email https://www.googleapis.com/auth/calendar"
    />
  );
}

export default LoginGoogle;
