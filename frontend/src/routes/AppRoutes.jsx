import {
  Routes,
  Route,
  Navigate,
} from "react-router-dom"

import Login from "../pages/Login"
import Register from "../pages/Register"
import AdminDashboard from "../pages/Admin/AdminDashboard"

// =========================
// PROTECTED ROUTE
// =========================

const ProtectedRoute = ({
  children,
  role,
}) => {

  const token = localStorage.getItem(
    "token"
  )

  const userRole = localStorage.getItem(
    "role"
  )

  // =========================
  // NOT LOGGED IN
  // =========================

  if (!token) {

    return <Navigate to="/" />
  }

  // =========================
  // ROLE CHECK
  // =========================

  if (
    role &&
    userRole !== role
  ) {

    return <Navigate to="/" />
  }

  return children
}

function AppRoutes() {

  return (

    <Routes>

      {/* ========================= */}
      {/* LOGIN */}
      {/* ========================= */}

      <Route
        path="/"
        element={<Login />}
      />

      {/* ========================= */}
      {/* REGISTER */}
      {/* ========================= */}

      <Route
        path="/register"
        element={<Register />}
      />

      {/* ========================= */}
      {/* ADMIN DASHBOARD */}
      {/* ========================= */}

      <Route
        path="/admin"
        element={
          <ProtectedRoute role="admin">

            <AdminDashboard />

          </ProtectedRoute>
        }
      />

    </Routes>
  )
}

export default AppRoutes