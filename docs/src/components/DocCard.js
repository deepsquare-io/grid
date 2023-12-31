import React from 'react';

import { Card, CardContent, Link, Typography } from '@mui/joy';

export default function DocCard({ children, title, href, startDecorator, sx }) {
  return (
    <Card
      orientation="horizontal"
      variant="outlined"
      sx={{
        p: '2rem',
        '&:hover': {
          boxShadow: 'md',
          borderColor: 'neutral.outlinedHoverBorder',
        },
        ...sx,
      }}
    >
      <CardContent>
        <Link overlay underline="none" href={href}>
          <Typography level="body1" startDecorator={startDecorator}>
            {title}
          </Typography>
        </Link>
        <Typography level="body2" sx={{ mt: 1 }}>
          {children}
        </Typography>
      </CardContent>
    </Card>
  );
}
