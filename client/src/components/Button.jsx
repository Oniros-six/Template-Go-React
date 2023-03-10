import React from "react";

const Button = (props) => {
    const {nombre, clases, icon} = props;
    return (
        <button className={clases} {...props}>
            {nombre}
            <i className={icon}/>
        </button>
        
    )

}

export default Button;