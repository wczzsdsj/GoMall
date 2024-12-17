package middleware

import (
	"context"

	frontendUtils "gomall/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

// type SessionUserIdKey string

// const SessionUserId SessionUserIdKey = "user_id"

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		// 把session写入到context里，写业务逻辑的地方就可以直接取context
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, s.Get("user_id"))
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		user_id := s.Get("user_id")
		if user_id == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
