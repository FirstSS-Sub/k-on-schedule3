import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';

// Route関連
import { Link } from 'react-router-dom'

const ResponsiveDrawerListItem = (to: string, onClick: () => void, icon: object, text: string) => (
    <ListItem button component={Link} to={to} onClick={onClick}>
        <ListItemIcon>
            {icon}
        </ListItemIcon>
        <ListItemText primary={text} />
    </ListItem>
);

export default ResponsiveDrawerListItem;
