import React, {useEffect, useState} from 'react';
import {
  AppBar,
  Avatar,
  Button,
  IconButton,
  makeStyles,
  Toolbar,
  Typography,
} from '@material-ui/core';
import {useHistory} from 'react-router-dom';
import {useDispatch, useSelector} from 'react-redux';
import {getMeRequest} from '../redux/ducks/userSlice';
import {Add, Menu} from '@material-ui/icons';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  loginButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
  createButton: {
    marginLeft: theme.spacing(2),
  },
}));

const Header = () => {
  const dispatch = useDispatch();
  const classes = useStyles();
  const history = useHistory();
  const {profile} = useSelector((state) => state.user);

  const [showCreate, setShowCreate] = useState(false);

  useEffect(() => {
    if (localStorage.getItem('token')) {
      dispatch(getMeRequest({token: localStorage.getItem('token') || ''}));
    }
  }, [dispatch]);

  return (
      <AppBar position="static">
        <Toolbar>
          <IconButton edge="start" className={classes.menuButton}
                      color="inherit" aria-label="menu"
                      onClick={() => history.push('/')}>
            <Menu/>
          </IconButton>
          <Typography variant="h6" className={classes.title}>
            Godutch
            {profile &&
            <Button variant="outlined" color="inherit"
                    className={classes.createButton}
                    onClick={() => setShowCreate(!showCreate)}
                    startIcon={<Add/>}>Create</Button>}
          </Typography>
          {profile ? (
                  <IconButton color="inherit"
                              onClick={() => history.push('/profile')}>
                    <Avatar>{profile.name.charAt(0)}</Avatar>
                  </IconButton>) :
              <>
                <Button className={classes.loginButton} color="inherit"
                        onClick={() => history.push('/login')}>Login</Button>
                <Button color="secondary" variant={'outlined'}
                        onClick={() => history.push('/signup')}>Signup</Button>
              </>
          }
        </Toolbar>
      </AppBar>
  );
};

export default Header;