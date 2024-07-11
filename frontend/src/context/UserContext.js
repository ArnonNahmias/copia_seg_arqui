import React, { createContext, useState, useEffect } from "react";

export const UserContext = createContext();

export const UserProvider = ({ children }) => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    // Aquí puedes agregar la lógica para obtener el usuario autenticado
    const fetchUser = async () => {
      // Supongamos que tienes una API para obtener el usuario actual
      const userData = await getUser(); // Reemplaza esto con tu lógica de autenticación
      setUser(userData);
    };

    fetchUser();
  }, []);

  return (
    <UserContext.Provider value={{ user, setUser }}>
      {children}
    </UserContext.Provider>
  );
};
