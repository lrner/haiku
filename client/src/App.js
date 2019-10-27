import React from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';


import Main from './components/Main'

function App() {
  return (
    <div className="App">
      <h1>
        A Haiku of Babel
      </h1>

      <Main />

      <footer>
        Inspired by <a href="https://en.wikipedia.org/wiki/The_Library_of_Babel">The Library of Babel</a> by Jorge Luis Borges and <a href="https://libraryofbabel.info/">libraryofbabel.info</a>
      </footer>
    </div>
  );
}

export default App;
