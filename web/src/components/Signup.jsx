import React, {useEffect, useState} from 'react';
import {
  AppBar,
  Button,
  CircularProgress,
  Grid,
  IconButton,
  makeStyles,
  Snackbar,
  TextField,
  Toolbar,
  Typography,
} from '@material-ui/core';
import {useDispatch, useSelector} from 'react-redux';
import {useHistory} from 'react-router-dom';
import {signupRequest} from '../redux/ducks/userSlice';
import CloseIcon from '@material-ui/icons/Close';

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
    flexWrap: 'wrap',
    '& > *': {
      margin: theme.spacing(1),
    },
  },
  title: {
    flexGrow: 1,
  },
}));

const Signup = () => {
  const classes = useStyles();
  const dispatch = useDispatch();
  const history = useHistory();
  const {profile, loading, error} = useSelector((state) => state.user);

  const [open, setOpen] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [name, setName] = useState('');

  useEffect(() => {
    if (error) {
      setOpen(true);
    }

    if (profile) {
      history.push('/login');
    }
  }, [dispatch, history, error, profile]);

  function handleSignup() {
    if (email && password && name) {
      dispatch(signupRequest({email: email, password: password, name: name}));

      setEmail('');
      setPassword('');
      setName('');
    }
  }

  const handleClose = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }

    setOpen(false);
  };

  return (
      <React.Fragment>
        <Grid item xs={false} sm={4}/>
        <Grid item xs={12} sm={4}>
          {loading ?
              <CircularProgress/> :
              (
                  <Grid className={classes.root} container direction={'column'}>
                    <Grid item>
                      <AppBar position="static">
                        <Toolbar>
                          <Typography variant="h6" className={classes.title}>
                            Signup
                          </Typography>
                        </Toolbar>
                      </AppBar>
                    </Grid>
                    <Grid item>
                      <form>
                        <TextField label={'name'} variant={'outlined'}
                                   margin="normal" fullWidth={true} required
                                   value={name}
                                   onChange={(e) => setName(e.target.value)}/>
                        <TextField label={'email'} type="email"
                                   variant={'outlined'}
                                   margin="normal" fullWidth={true} required
                                   value={email}
                                   onChange={(e) => setEmail(e.target.value)}/>
                        <TextField label={'password'} type="password"
                                   variant={'outlined'}
                                   margin="normal" fullWidth={true} required
                                   value={password}
                                   onChange={(e) => setPassword(e.target.value)}
                                   autoComplete="current-password"/>
                        <Button color={'primary'} variant={'contained'}
                                size={'large'}
                                fullWidth={true}
                                onClick={() => handleSignup()}>Submit</Button>
                      </form>
                    </Grid>
                  </Grid>
              )}
        </Grid>
        <Grid item xs={false} sm={4}/>

        <Snackbar
            anchorOrigin={{
              vertical: 'bottom',
              horizontal: 'left',
            }}
            open={open}
            autoHideDuration={3000}
            onClose={handleClose}
            message={error}
            action={
              <React.Fragment>
                <IconButton size="small" aria-label="close" color="inherit"
                            onClick={handleClose}>
                  <CloseIcon fontSize="small"/>
                </IconButton>
              </React.Fragment>
            }
        />
      </React.Fragment>
  );
};

export default Signup;
