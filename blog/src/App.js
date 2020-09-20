import React from 'react';
import { Router } from '@reach/router';
import './App.css';

//Routes
import Home from './Pages/Home.jsx';
import Contact from './Pages/Contact.jsx';
import Home from './Pages/Home.jsx';


function App() {
  const navLinks = [
    {
      text: 'Contact',
      path: '/contact',
      icon: 'ion-ios-megaphone'
    },
    {
      text: 'Blog',
      path: '/blog',
      icon: 'ion-ios=bonfire'
    },
    {
      text: 'Portfolio',
      path: '/portfolio',
      icon: 'ion-ios-briefcase'
    }
  ]


  return (
    <div className="App">
      <Router>

      </Router>
    </div>
  );
}

export default App;
