import {takeLatest} from 'redux-saga/effects';
import {
  handleGetMe,
  handleLogin,
  handleLogout,
  handleSignup,
} from './handlers/user';
import {
  getMeRequest,
  loginRequest,
  logoutRequest,
  signupRequest,
} from '../ducks/userSlice';

export function* watcherSaga() {
  yield takeLatest(loginRequest.type, handleLogin);
  yield takeLatest(signupRequest.type, handleSignup);
  yield takeLatest(getMeRequest.type, handleGetMe);
  yield takeLatest(logoutRequest.type, handleLogout);
}