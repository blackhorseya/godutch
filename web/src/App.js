import React from 'react';
import {Route, Switch} from 'react-router-dom';
import {Grid} from '@material-ui/core';
import Header from './components/Header';
import Dashboard from './components/Dashboard';
import Profile from './components/Profile';
import Login from './components/Login';
import Signup from './components/Signup';

function App() {
  return (
      <Grid container direction={'column'}>
        <Grid item>
          <Header/>
        </Grid>
        <Grid item container>
          <Switch>
            <Route exact from="/">
              <Dashboard/>
            </Route>
            <Route from="/profile">
              <Profile/>
            </Route>
            <Route from="/login">
              <Login/>
            </Route>
            <Route from="/signup">
              <Signup/>
            </Route>
          </Switch>
        </Grid>
      </Grid>
  );
}

export default App;
