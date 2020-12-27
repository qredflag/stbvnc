package main

type fb_bitfield struct {
	offset    uint32 /* beginning of bitfield	*/
	length    uint32 /* length of bitfield		*/
	msb_right uint32 /* != 0 : Most significant bit is */
}

type fb_var_screeninfo struct {
	xres           uint32 /* visible resolution		*/
	yres           uint32
	xres_virtual   uint32 /* virtual resolution		*/
	yres_virtual   uint32
	xoffset        uint32 /* offset from virtual to visible */
	yoffset        uint32 /* resolution			*/
	bits_per_pixel uint32
	grayscale      uint32

	/* >1 = FOURCC			*/
	red    fb_bitfield
	green  fb_bitfield
	blue   fb_bitfield
	transp fb_bitfield

	nonstd   uint32 /* != 0 Non standard pixel format */
	activate uint32 /* see FB_ACTIVATE_*		*/
	height   uint32 /* height of picture in mm    */
	width    uint32 /* width of picture in mm     */

	accel_flags uint32 /* (OBSOLETE) see fb_info.flags */

	/* Timing: All values in pixclocks, except pixclock (of course) */
	pixclock     uint32 /* pixel clock in ps (pico seconds) */
	left_margin  uint32 /* time from sync to picture	*/
	right_margin uint32 /* time from picture to sync	*/
	upper_margin uint32 /* time from sync to picture	*/
	lower_margin uint32
	hsync_len    uint32    /* length of horizontal sync	*/
	vsync_len    uint32    /* length of vertical sync	*/
	sync         uint32    /* see FB_SYNC_*		*/
	vmode        uint32    /* see FB_VMODE_*		*/
	rotate       uint32    /* angle we rotate counter clockwise */
	colorspace   uint32    /* colorspace for FOURCC-based modes */
	reserved     [4]uint32 /* Reserved for future compatibility */
}
