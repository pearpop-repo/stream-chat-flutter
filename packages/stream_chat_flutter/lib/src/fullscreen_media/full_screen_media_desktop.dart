import 'package:cached_network_image/cached_network_image.dart';
import 'package:dart_vlc/dart_vlc.dart';
import 'package:flutter/material.dart';
import 'package:native_context_menu/native_context_menu.dart';
import 'package:photo_view/photo_view.dart';
import 'package:stream_chat_flutter/src/context_menu_items/download_menu_item.dart';
import 'package:stream_chat_flutter/src/extension.dart';
import 'package:stream_chat_flutter/src/platform_widgets/platform_widget_builder.dart';
import 'package:stream_chat_flutter/stream_chat_flutter.dart';

/// Return action for coming back from pages
enum ReturnActionType {
  /// No return action
  none,

  /// Go to reply message action
  reply,
}

/// Callback when show message is tapped
typedef ShowMessageCallback = void Function(Message message, Channel channel);

/// A full screen image widget
class FullScreenMediaDesktop extends StatefulWidget {
  /// Instantiate a new FullScreenImage
  const FullScreenMediaDesktop({
    Key? key,
    required this.mediaAttachments,
    required this.message,
    this.startIndex = 0,
    String? userName,
    this.onShowMessage,
    this.attachmentActionsModalBuilder,
    this.autoplayVideos = false,
  })  : userName = userName ?? '',
        super(key: key);

  /// The url of the image
  final List<Attachment> mediaAttachments;

  /// Message where attachments are attached
  final Message message;

  /// First index of media shown
  final int startIndex;

  /// Username of sender
  final String userName;

  /// Callback for when show message is tapped
  final ShowMessageCallback? onShowMessage;

  /// Widget builder for attachment actions modal
  /// [defaultActionsModal] is the default [AttachmentActionsModal] config
  /// Use [defaultActionsModal.copyWith] to easily customize it
  final AttachmentActionsBuilder? attachmentActionsModalBuilder;

  /// Auto-play videos when page is opened
  final bool autoplayVideos;

  @override
  _FullScreenMediaDesktopState createState() => _FullScreenMediaDesktopState();
}

class _FullScreenMediaDesktopState extends State<FullScreenMediaDesktop>
    with SingleTickerProviderStateMixin {
  late final AnimationController _animationController;
  late final PageController _pageController;

  late final _curvedAnimation = CurvedAnimation(
    parent: _animationController,
    curve: Curves.easeOut,
    reverseCurve: Curves.easeIn,
  );

  final _opacityTween = Tween<double>(begin: 1, end: 0);
  late final _opacityAnimation = _opacityTween.animate(
    CurvedAnimation(
      parent: _animationController,
      curve: const Interval(0, 0.6, curve: Curves.easeOut),
    ),
  );

  late final ValueNotifier<int> _currentPage = ValueNotifier(widget.startIndex);

  final videoPackages = <String, DesktopVideoPackage>{};

  @override
  void initState() {
    super.initState();
    _animationController = AnimationController(
      vsync: this,
      duration: const Duration(milliseconds: 300),
    );
    _pageController = PageController(initialPage: widget.startIndex);
    for (var i = 0; i < widget.mediaAttachments.length; i++) {
      final attachment = widget.mediaAttachments[i];
      if (attachment.type != 'video') continue;
      final package = DesktopVideoPackage(attachment);
      videoPackages[attachment.id] = package;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      resizeToAvoidBottomInset: false,
      body: Stack(
        children: [
          PageView.builder(
            controller: _pageController,
            onPageChanged: (val) {
              _currentPage.value = val;

              if (videoPackages.isEmpty) {
                return;
              }

              final currentAttachment = widget.mediaAttachments[val];

              for (final e in videoPackages.values) {
                if (e.attachment != currentAttachment) {
                  e.player.pause();
                }
              }

              if (widget.autoplayVideos && currentAttachment.type == 'video') {
                final package = videoPackages[currentAttachment.id]!;
                package.player.play();
              }
            },
            itemBuilder: (context, index) {
              final attachment = widget.mediaAttachments[index];
              if (attachment.type == 'image' || attachment.type == 'giphy') {
                final imageUrl = attachment.imageUrl ??
                    attachment.assetUrl ??
                    attachment.thumbUrl;
                return AnimatedBuilder(
                  animation: _curvedAnimation,
                  builder: (context, child) => ContextMenuRegion(
                    onItemSelected: (item) {
                      item.onSelected?.call();
                    },
                    menuItems: [
                      DownloadMenuItem(
                        attachment: attachment,
                      ),
                    ],
                    child: PhotoView(
                      loadingBuilder: (context, image) => const Offstage(),
                      imageProvider: (imageUrl == null &&
                              attachment.localUri != null &&
                              attachment.file?.bytes != null)
                          ? Image.memory(attachment.file!.bytes!).image
                          : CachedNetworkImageProvider(imageUrl!),
                      maxScale: PhotoViewComputedScale.covered,
                      minScale: PhotoViewComputedScale.contained,
                      heroAttributes: PhotoViewHeroAttributes(
                        tag: widget.mediaAttachments,
                      ),
                      backgroundDecoration: BoxDecoration(
                        color: ColorTween(
                          begin: ChannelHeaderTheme.of(context).color,
                          end: Colors.black,
                        ).lerp(_curvedAnimation.value),
                      ),
                      onTapUp: (a, b, c) {
                        if (_animationController.isCompleted) {
                          _animationController.reverse();
                        } else {
                          _animationController.forward();
                        }
                      },
                    ),
                  ),
                );
              } else if (attachment.type == 'video') {
                final package = videoPackages[attachment.id]!;
                package.player.open(
                  Playlist(
                    medias: [
                      Media.network(package.attachment.assetUrl),
                    ],
                  ),
                  autoStart: widget.autoplayVideos,
                );
                return InkWell(
                  splashFactory: NoSplash.splashFactory,
                  onTap: () {
                    if (_animationController.isCompleted) {
                      _animationController.reverse();
                    } else {
                      _animationController.forward();
                    }
                  },
                  child: Padding(
                    padding: const EdgeInsets.symmetric(
                      vertical: 50,
                    ),
                    child: Video(
                      player: package.player,
                    ),
                  ),
                );
              }
              return const SizedBox();
            },
            itemCount: widget.mediaAttachments.length,
          ),
          FadeTransition(
            opacity: _opacityAnimation,
            child: ValueListenableBuilder<int>(
              valueListenable: _currentPage,
              builder: (context, value, child) => Column(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  GalleryHeader(
                    userName: widget.userName,
                    sentAt: context.translations.sentAtText(
                      date: widget.message.createdAt,
                      time: widget.message.createdAt,
                    ),
                    onBackPressed: () {
                      Navigator.of(context).pop();
                    },
                    message: widget.message,
                    currentIndex: value,
                    onShowMessage: () {
                      widget.onShowMessage?.call(
                        widget.message,
                        StreamChannel.of(context).channel,
                      );
                    },
                    attachmentActionsModalBuilder:
                        widget.attachmentActionsModalBuilder,
                  ),
                  if (!widget.message.isEphemeral)
                    GalleryFooter(
                      currentPage: value,
                      totalPages: widget.mediaAttachments.length,
                      mediaAttachments: widget.mediaAttachments,
                      message: widget.message,
                      mediaSelectedCallBack: (val) {
                        _currentPage.value = val;
                        _pageController.animateToPage(
                          val,
                          duration: const Duration(milliseconds: 300),
                          curve: Curves.easeInOut,
                        );
                        Navigator.pop(context);
                      },
                    ),
                ],
              ),
            ),
          ),
          if (widget.mediaAttachments.length > 1) ...[
            if (_currentPage.value != widget.mediaAttachments.length - 1) ...[
              GalleryNavigationItem(
                icon: Icons.chevron_right,
                right: 0,
                opacityAnimation: _opacityAnimation,
                currentPage: _currentPage,
                onClick: () {
                  _pageController.nextPage(
                    duration: const Duration(milliseconds: 350),
                    curve: Curves.easeIn,
                  );
                  setState(() => _currentPage.value++);
                },
              ),
            ],
            if (_currentPage.value != 0) ...[
              GalleryNavigationItem(
                icon: Icons.chevron_left,
                left: 0,
                opacityAnimation: _opacityAnimation,
                currentPage: _currentPage,
                onClick: () {
                  _pageController.previousPage(
                    duration: const Duration(milliseconds: 350),
                    curve: Curves.easeOut,
                  );
                  setState(() => _currentPage.value--);
                },
              ),
            ],
          ],
        ],
      ),
    );
  }

  @override
  void dispose() {
    _animationController.dispose();
    _pageController.dispose();
    super.dispose();
  }
}

/// A widget for desktop and web users to be able to navigate left and right
/// through a gallery of images.
class GalleryNavigationItem extends StatelessWidget {
  /// Builds a [GalleryNavigationItem].
  const GalleryNavigationItem({
    Key? key,
    required this.icon,
    required this.onClick,
    required this.opacityAnimation,
    required this.currentPage,
    this.left,
    this.right,
  }) : super(key: key);

  /// The icon to display.
  final IconData icon;

  /// The callback to perform when the button is clicked.
  final VoidCallback onClick;

  /// The animation for showing & hiding this widget.
  final Animation<double> opacityAnimation;

  /// The value to use for .
  final ValueNotifier<int> currentPage;

  /// The left-hand placement of the button.
  final double? left;

  /// The right-hand placement of the button.
  final double? right;

  @override
  Widget build(BuildContext context) => PlatformWidgetBuilder(
        desktop: (_, child) => child,
        web: (_, child) => child,
        child: Positioned(
          left: left,
          right: right,
          top: MediaQuery.of(context).size.height / 2,
          child: FadeTransition(
            opacity: opacityAnimation,
            child: ValueListenableBuilder<int>(
              valueListenable: currentPage,
              builder: (context, value, child) => GestureDetector(
                onTap: onClick,
                child: Icon(
                  icon,
                  size: 50,
                ),
              ),
              child: GestureDetector(
                onTap: onClick,
                child: Icon(
                  icon,
                  size: 50,
                ),
              ),
            ),
          ),
        ),
      );
}

/// Class for packaging up things required for videos
class DesktopVideoPackage {
  /// Constructor for creating [VideoPackage]
  DesktopVideoPackage(
    this.attachment, {
    this.showControls = true,
  }) : player = Player(
          id: int.parse(
            attachment.id.characters
                .getRange(0, 10)
                .toString()
                .replaceAll(RegExp('[^0-9]'), ''),
          ),
        );

  final Attachment attachment;
  final Player player;
  final bool showControls;
}
