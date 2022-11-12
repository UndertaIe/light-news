import React from 'react'
import Search from './pages/Search'
import { Router } from '@reach/router'

const App: React.FC = (props) => {
  return (
    <>
      <Router>
        <Search path="/search"/>
      </Router>
    </>
  )
}

export default App