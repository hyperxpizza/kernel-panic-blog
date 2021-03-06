import React from 'react';
import {Route, BrowserRouter as Router} from 'react-router-dom';

import Home from './pages/Home';
import Login from './pages/Login';
import Navigation from './components/Navigation/Navigation.js';
import Post from './pages/Post';
import Admin from './pages/Admin';

function App() {
  return (
    <Router>
      <Navigation />
      <Route path="/" exact component={Home} />
      <Route path="/login" exact component={Login} />
      <Route path="/admin" exact component={Admin} />
      <Route path="/post/:slug" exact component={Post} />
    </Router>
  );
}

export default App;
